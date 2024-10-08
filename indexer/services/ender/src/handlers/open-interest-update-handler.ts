import { logger } from '@nemo-network-indexer/base/build';
import {
  PerpetualMarketFromDatabase,
  PerpetualMarketModel,
  PerpetualMarketColumns,
  MarketMessageContents,
  TradingMarketMessageContents,
  TradingPerpetualMarketMessage,
} from '@nemo-network-indexer/postgres/build/src';
import { OpenInterestUpdateEventV1 } from '@nemo-network-indexer/v4-protos';
import _ from 'lodash';
import * as pg from 'pg';

import { ConsolidatedKafkaEvent } from '../lib/types';
import { Handler } from './handler';

export class OpenInterestUpdateHandler extends Handler<OpenInterestUpdateEventV1> {
  eventType: string = 'OpenInterestUpdate';

  public getParallelizationIds(): string[] {
    return [];
  }

  // eslint-disable-next-line @typescript-eslint/require-await
  public async internalHandle(resultRow: pg.QueryResultRow): Promise<ConsolidatedKafkaEvent[]> {
    const perpetualMarkets: PerpetualMarketFromDatabase[] = resultRow.open_interest_updates.map(
      (openInterestUpdate: PerpetualMarketFromDatabase) => {
        return PerpetualMarketModel.fromJson(
          openInterestUpdate as object) as PerpetualMarketFromDatabase;
      },
    );
    logger.info({
      at: 'OpenInterestUpdateHandler#handle',
      message: 'Received OpenInterestUpdate',
    });

    if (perpetualMarkets.length === 0) {
      return [];
    }

    return [
      this.generateConsolidatedMarketKafkaEvent(
        JSON.stringify(generateMarketMessage(perpetualMarkets)),
      ),
    ];
  }
}

function generateMarketMessage(
  perpetualMarkets: PerpetualMarketFromDatabase[],
): MarketMessageContents {
  const tradingMarketMessageContents: TradingMarketMessageContents = _.chain(perpetualMarkets)
    .keyBy(PerpetualMarketColumns.ticker)
    .mapValues((perpetualMarket: PerpetualMarketFromDatabase): TradingPerpetualMarketMessage => {
      return {
        id: perpetualMarket.id,
        baseOpenInterest: perpetualMarket.baseOpenInterest,
      };
    })
    .value();

  return {
    trading: tradingMarketMessageContents,
  };
}
