/* @flow */
// JSX complication requires a reference to `React` and flow doesn't understand
// `import React, { Component, PropTypes } from 'react';`.
import React from 'react';
const { Component, PropTypes } = React;
import d3 from 'd3';
import { LineChart } from 'react-d3-components';

import { formatDateTime, formatDate, formatTime } from './date-utils';
import { ChannelShape, DatapointShape } from './prop-types';

export class DatapointChart extends Component {
  formatTick(d: Date): string {
    const hours = d.getHours();
    if (hours === 0) {
      return formatDate(d);
    } else {
      return formatTime(d);
    }
  }

  render(): any {
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
    const xAxis = { tickFormat: this.formatTick.bind(this) };
    const yAxis = { label: 'Deg C' };

    const tooltipHtml = (_, data) => {
      const date = formatDateTime(data.x);
      return `${data.y} at ${date}`;
    };

    return (
      <LineChart data={data}
                 width={width}
                 height={height}
                 margin={margin}
                 tooltipHtml={tooltipHtml}
                 xAxis={xAxis}
                 yAxis={yAxis}
                 xScale={xScale} />
    );
  }
}
// TODO: Use `static get propTypes()` as soon as flow supports static getter.
DatapointChart.propTypes = {
  datapoints: PropTypes.arrayOf(DatapointShape)
};

export class Channel extends Component {
  render(): any {
    const channel = this.props.channel;
    const url = `/channels/${channel.id}/datapoints`;
    const latest = channel.datapoints[channel.datapoints.length - 1];
    const date = formatDateTime(latest.at);
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
Channel.propTypes = {
  channel: ChannelShape.isRequired
};

export class ChannelList extends Component {
  render(): any {
    if (!this.props.channels) {
      return null;
    }

    const listItems = this.props.channels.map((channel) => <Channel channel={channel} key={channel.id} />);
    return <div className="channel-list">{listItems}</div>;
  }
}
ChannelList.propTypes = {
  channels: PropTypes.arrayOf(ChannelShape).isRequired
};

export class ErrorMessage extends Component {
  render(): any {
    if (!this.props.error) {
      return null;
    }
    return <p className="error-message">Failed to fetch data: {this.props.error.toString()}</p>;
  }
}
ErrorMessage.propTypes = {
  channels: PropTypes.arrayOf(ChannelShape),
};

export class Dashboard extends Component {
  render(): any {
    return (
      <div className="dashboard">
        <h1>Beerserver Dashboard</h1>
        <ErrorMessage error={this.props.error} />
        <ChannelList channels={this.props.channels} />
      </div>
    );
  }
}
