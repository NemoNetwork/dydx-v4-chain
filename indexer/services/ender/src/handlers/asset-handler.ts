import {
  AssetFromDatabase,
  AssetModel,
  assetRefresher,
} from '@nemo-network-indexer/postgres/build/src';
import { AssetCreateEventV1 } from '@nemo-network-indexer/v4-protos';
import * as pg from 'pg';

import { ConsolidatedKafkaEvent } from '../lib/types';
import { Handler } from './handler';

export class AssetCreationHandler extends Handler<AssetCreateEventV1> {
  eventType: string = 'AssetCreateEvent';

  public getParallelizationIds(): string[] {
    return [];
  }

  // eslint-disable-next-line @typescript-eslint/require-await
  public async internalHandle(resultRow: pg.QueryResultRow): Promise<ConsolidatedKafkaEvent[]> {
    const asset: AssetFromDatabase = AssetModel.fromJson(
      resultRow.asset) as AssetFromDatabase;
    assetRefresher.addAsset(asset);
    return [];
  }
}
