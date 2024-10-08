import {
  IsolationLevel,
  Transaction,
  assetRefresher,
  perpetualMarketRefresher,
} from '@nemo-network-indexer/postgres/build/src';

export async function refreshDataCaches(): Promise<void> {
  const txId: number = await Transaction.start();
  await Transaction.setIsolationLevel(txId, IsolationLevel.READ_COMMITTED);

  await Promise.all([
    perpetualMarketRefresher.updatePerpetualMarkets({ txId, readReplica: true }),
    assetRefresher.updateAssets({ txId, readReplica: true }),
  ]);

  await Transaction.rollback(txId);
}
