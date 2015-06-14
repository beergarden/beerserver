type Datapoint = {
  at: Date;
  value: number;
};

type Channel = {
  id: string;
  name: string;
  datapoints: Array<Datapoint>;
};
