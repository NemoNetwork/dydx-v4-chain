/**
 * Example of an auxo event json:
{
  "upgrade_tag": "4aea7a61",
  "prefix": "dev4",
  "region": "ap-northeast-1",
  "regionAbbrev": "apne1"
}
 */
export interface AuxoEventJson {
  upgrade_tag: string;
  prefix: string;
  region: string;
  // In our naming we often times use the appreviated region name
  regionAbbrev: string;
<<<<<<< HEAD
  onlyRunDbMigrationAndCreateKafkaTopics: boolean;
=======
  addNewKafkaTopics: boolean;
>>>>>>> d36863c7 (Adam/ct 861 websocket topic for blockheight (#1705))
}

// EcsServiceName to task definition arn mapping
export type TaskDefinitionArnMap = _.Dictionary<string>;

export enum EcsServiceNames {
  COMLINK = 'comlink',
  ENDER = 'ender',
  ROUNDTABLE = 'roundtable',
  SOCKS = 'socks',
  VULCAN = 'vulcan',
}
