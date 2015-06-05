import 'whatwg-fetch';
import React from 'react';

import { Dashboard } from './components';

fetchChannels()
  .then(renderChannels);

function fetchChannels() {
  return fetch('/channels')
    .then((response) => response.json())
    .then((channels) => Promise.all(channels.map(fetchDatapoints)));
}

function fetchDatapoints(channel) {
  return fetch(`/channels/${channel.id}/datapoints`)
    .then((response) => response.json())
    .then((datapoints) => {
      channel.datapoints = datapoints.map((d) => {
        return { at: new Date(d.at), value: d.value };
      });
      return channel;
    });
}

function renderChannels(channels) {
  React.render(
    <Dashboard channels={channels} />,
    document.getElementById('world')
  );
}
