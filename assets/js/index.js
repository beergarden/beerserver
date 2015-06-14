/* @flow */
import _ from 'whatwg-fetch';
import React from 'react';

import { Dashboard } from './components';

fetchChannels()
  .then(renderDashboard)
  .catch(renderError);

function fetchChannels(): Promise<Array<Channel>> {
  return fetch('/channels')
    .then((response: Response) => {
      if (response.ok) {
        return response.json();
      } else {
        return response.text().then((t: string) => Promise.reject(new Error(t)));
      }
    })
    .then((channels) => Promise.all(channels.map(fetchDatapoints)));
}

function fetchDatapoints(channel: Channel): Promise<Channel> {
  return fetch(`/channels/${channel.id}/datapoints`)
    .then((response: Response) => response.json())
    .then((datapoints) => {
      channel.datapoints = datapoints.map((d) => {
        return { at: new Date(d.at), value: d.value };
      });
      return channel;
    });
}

function renderDashboard(channels: Array<Channel>) {
  React.render(
    <Dashboard channels={channels} />,
    document.getElementById('world')
  );
}

function renderError(error) {
  React.render(
    <Dashboard error={error} />,
    document.getElementById('world')
  );
}
