export type Datapoint = {
  at: Date;
  value: number;
};

export type Channel = {
  id: string;
  name: string;
  datapoints: Array<Datapoint>;
};
