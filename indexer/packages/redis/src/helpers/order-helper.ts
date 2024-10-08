import { OrderTable, PerpetualMarketFromDatabase, protocolTranslations } from '@nemo-network-indexer/postgres/build/src';
import { IndexerOrder, RedisOrder, RedisOrder_TickerType } from '@nemo-network-indexer/v4-protos';

/**
 * Creates a `RedisOrder` given an `Order` and the corresponding `PerpetualMarket` for the `Order`.
 * @param order
 * @param perpetualMarket
 * @returns
 */
export function convertToRedisOrder(
  order: IndexerOrder,
  perpetualMarket: PerpetualMarketFromDatabase,
): RedisOrder {
  return {
    order,
    id: OrderTable.orderIdToUuid(order.orderId!),
    ticker: perpetualMarket.ticker,
    tickerType: RedisOrder_TickerType.TICKER_TYPE_PERPETUAL,
    price: protocolTranslations.subticksToPrice(
      order.subticks.toString(),
      perpetualMarket,
    ),
    size: protocolTranslations.quantumsToHumanFixedString(
      order.quantums.toString(),
      perpetualMarket.atomicResolution,
    ),
  };
}
