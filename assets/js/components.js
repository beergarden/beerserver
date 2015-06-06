import React from 'react';
import d3 from 'd3';
import { LineChart } from 'react-d3-components';

export class DatapointChart {
  render() {
    const values = this.props.datapoints.map((d) => {
      return { x: d.at, y: d.value };
    });
    const data = { label: '', values };

    const width = 800;
    const height = 200;
    const margin = { top: 10, bottom: 20, left: 50, right: 20 };

    const xs = values.map((v) => v.x);
    const minX = Math.min(...xs);
    const maxX = Math.max(...xs);

    const xScale = d3.time.scale()
      .domain([minX, maxX])
      .range([0, width - margin.left - margin.right]);

    return (
      <LineChart data={data}
                 width={width}
                 height={height}
                 margin={margin}
                 xScale={xScale} />
    );
  }
}

export class Channel extends React.Component {
  formatDate(date) {
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    const hours = date.getHours();
    const minutes = date.getMinutes();
    return `${year}/${month}/${day} ${hours}:${minutes}`;
  }

  render() {
    const channel = this.props.channel;
    const url = `/channels/${channel.id}/datapoints`;
    const latest = channel.datapoints[channel.datapoints.length - 1];
    const date = this.formatDate(latest.at);
    return (
      <div className="channel">
        <h2 className="name">{channel.name}</h2>
        <p>
          <span className="latest">Latest: {latest.value} at {date}</span>
          <span> </span>
          <span className="json"><a href={url}>JSON</a></span>
        </p>
        <DatapointChart datapoints={channel.datapoints} />
      </div>
    );
  }
}

export class ChannelList extends React.Component {
  render() {
    if (!this.props.channels) {
      return null;
    }

    const listItems = this.props.channels.map((channel) => <Channel channel={channel} key={channel.id} />);
    return <div className="channel-list">{listItems}</div>;
  }
}

export class ErrorMessage extends React.Component {
  render() {
    if (!this.props.error) {
      return null;
    }
    return <p className="error-message">Failed to fetch data: {this.props.error.toString()}</p>;
  }
}

export class Dashboard extends React.Component {
  render() {
    return (
      <div className="dashboard">
        <h1>Beerserver Dashboard</h1>
        <ErrorMessage error={this.props.error} />
        <ChannelList channels={this.props.channels} />
      </div>
    );
  }
}
