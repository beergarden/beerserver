/* @flow */
/* global require */
require('whatwg-fetch');
import React from 'react';

import { Dashboard } from './components';
import type { Channel } from './models';

(async () => {
  try {
    var channels = await fetchChannels();
    renderDashboard(channels);

    channels.forEach(async (channel) => {
      await fetchDatapoints(channel);
      renderDashboard(channels);
    });
  } catch (e) {
    renderError(e);
  }
})();

async function fetchChannels(): Promise<Array<Channel>> {
  var response = await fetch('/channels');
  if (!response.ok) {
    var errorMessage = await response.text();
    throw new Error(errorMessage);
  }
  return await response.json();
}

async function fetchDatapoints(channel: Channel): Promise<Channel> {
  var response = await fetch(`/channels/${channel.id}/datapoints`);
  var datapoints = await response.json();
  channel.datapoints = datapoints.map((d) => {
    return { at: new Date(d.at), value: d.value };
  });
  return channel;
}

function renderDashboard(channels: Array<Channel>) {
  React.render(
    <Dashboard channels={channels} />,
    document.getElementById('world')
  );
}

function renderError(error: Error) {
  React.render(
    <Dashboard error={error} />,
    document.getElementById('world')
  );
}
