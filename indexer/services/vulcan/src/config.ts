/**
 * Environment variables required by Vulcan.
 */

import {
  parseNumber,
  parseSchema,
  baseConfigSchema,
  parseBoolean,
  parseString,
} from '@nemo-network-indexer/base/build';
import {
  kafkaConfigSchema,
} from '@nemo-network-indexer/kafka/build/src';
import {
  redisConfigSchema,
} from '@nemo-network-indexer/redis/build/redis/src';

export const configSchema = {
  ...baseConfigSchema,
  ...kafkaConfigSchema,
  ...redisConfigSchema,

  BATCH_PROCESSING_ENABLED: parseBoolean({ default: true }),
  PROCESS_FROM_BEGINNING: parseBoolean({ default: false }),
  KAFKA_BATCH_PROCESSING_COMMIT_FREQUENCY_MS: parseNumber({
    default: 3_000,
  }),
  FLUSH_KAFKA_MESSAGES_INTERVAL_MS: parseNumber({
    default: 10,
  }),
  MAX_WEBSOCKET_MESSAGES_TO_QUEUE_PER_TOPIC: parseNumber({
    default: 20,
  }),
  // Set this flag to false during fast sync.
  SEND_WEBSOCKET_MESSAGES: parseBoolean({
    default: true,
  }),
  SEND_SUBACCOUNT_WEBSOCKET_MESSAGE_FOR_STATEFUL_ORDERS: parseBoolean({
    default: true,
  }),
  SEND_SUBACCOUNT_WEBSOCKET_MESSAGE_FOR_CANCELS_MISSING_ORDERS: parseBoolean({
    default: true,
  }),
  KAFKA_BROKER_URLS: parseString({
    default: 'kafka:9092',
  }),
};

export default parseSchema(configSchema);
