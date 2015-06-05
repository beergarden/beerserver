import 'whatwg-fetch';
import React from 'react';

class ChannelList extends React.Component {
  render() {
    const listItems = this.props.channels.map((channel) => {
      const url = `/channels/${channel.id}/datapoints`;
      return <li key={channel.id}><a href={url}>{channel.name}</a></li>;
    });
    return (
      <ul>{listItems}</ul>
    );
  }
}

class Dashboard extends React.Component {
  render() {
    return (
      <div>
        <h1>Beerserver Dashboard</h1>
        <ChannelList channels={this.props.channels} />
      </div>
    );
  }
}

fetch('/channels')
  .then((response) => {
    return response.json();
  })
  .then((channels) => {
    React.render(
      <Dashboard channels={channels} />,
      document.getElementById('world')
    );
  });
