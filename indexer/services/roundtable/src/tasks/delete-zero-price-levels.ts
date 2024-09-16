import {
  stats,
} from '@nemo-network-indexer/base/src';
import {
  OrderSide,
  PerpetualMarketFromDatabase,
  PerpetualMarketTable,
} from '@nemo-network-indexer/postgres';
import { OrderbookLevels, OrderbookLevelsCache, PriceLevel } from '@nemo-network-indexer/redis/src';
import _ from 'lodash';

import config from '../config';
import { redisClient } from '../helpers/redis';

export default async function runTask(): Promise<void> {
  let numDeletedZeroPriceLevels: number = 0;

  const perpMarkets: PerpetualMarketFromDatabase[] = await PerpetualMarketTable.findAll({}, []);
  for (const perpetualMarket of perpMarkets) {
    const priceLevels: OrderbookLevels = await OrderbookLevelsCache.getOrderBookLevels(
      perpetualMarket.ticker,
      redisClient,
      {
        removeZeros: false,
      },
    );

    const zeroBidLevels: PriceLevel[] = priceLevels.bids.filter(
      (level: PriceLevel): boolean => level.quantums === '0',
    );
    const zeroAskLevels: PriceLevel[] = priceLevels.asks.filter(
      (level: PriceLevel): boolean => level.quantums === '0',
    );

    const removedLevels: boolean [] = await Promise.all(_.flatten([
      _.map(zeroBidLevels, (zeroBidLevel: PriceLevel): Promise<boolean> => {
        return OrderbookLevelsCache.deleteZeroPriceLevel({
          ticker: perpetualMarket.ticker,
          side: OrderSide.BUY,
          humanPrice: zeroBidLevel.humanPrice,
          client: redisClient,
        });
      }),
      _.map(zeroAskLevels, (zeroAskLevel: PriceLevel): Promise<boolean> => {
        return OrderbookLevelsCache.deleteZeroPriceLevel({
          ticker: perpetualMarket.ticker,
          side: OrderSide.SELL,
          humanPrice: zeroAskLevel.humanPrice,
          client: redisClient,
        });
      }),
    ]));

    numDeletedZeroPriceLevels += _.filter(
      removedLevels,
      (removed: boolean): boolean => removed === true,
    ).length;
  }

  stats.gauge(
    `${config.SERVICE_NAME}.delete_zero_price_levels.num_levels_deleted`,
    numDeletedZeroPriceLevels,
  );
}
