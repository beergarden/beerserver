/* @flow */
// JSX complication requires a reference to `React` and flow doesn't understand
// `import React, { Component, PropTypes } from 'react';`.
import React from 'react';
var { Component, PropTypes } = React;
import d3 from 'd3';
import { LineChart } from 'react-d3-components';

import { formatDateTime, formatDate, formatTime } from './date-utils';
import { ChannelShape, DatapointShape } from './prop-types';

export class DatapointChart extends Component {
  formatTick(d: Date): string {
    var hours = d.getHours();
    if (hours === 0) {
      return formatDate(d);
    } else {
      return formatTime(d);
    }
  }

  render(): any {
    var values = this.props.datapoints.map((d) => {
      return { x: d.at, y: d.value };
    });
    var data = { label: '', values };

    var width = 800;
    var height = 200;
    var margin = { top: 10, bottom: 20, left: 50, right: 20 };

    var xs = values.map((v) => v.x);
    var minX = Math.min(...xs);
    var maxX = Math.max(...xs);

    var xScale = d3.time.scale()
      .domain([minX, maxX])
      .range([0, width - margin.left - margin.right]);
    var xAxis = { tickFormat: this.formatTick.bind(this) };
    var yAxis = { label: 'Deg C' };

    var tooltipHtml = (_, d) => {
      var date = formatDateTime(d.x);
      return `${d.y} at ${date}`;
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
    var channel = this.props.channel;
    var url = `/channels/${channel.id}/datapoints`;
    var content;

    if (channel.datapoints) {
      if (channel.datapoints.length > 0) {
        var latest = channel.datapoints[channel.datapoints.length - 1];
        var date = formatDateTime(latest.at);
        content = (
          <div>
            <p>
              <span className="latest">Latest: {latest.value} at {date}</span>
              <span> </span>
              <span className="json"><a href={url}>JSON</a></span>
            </p>
            <DatapointChart datapoints={channel.datapoints} />
          </div>
        );
      } else {
        content = <p>No data</p>;
      }
    } else {
      content = <p className="loading">Loading...</p>;
    }
    return (
      <div className="channel">
        <h2 className="name">{channel.name}</h2>
        {content}
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

    var listItems = this.props.channels.map((channel) => <Channel channel={channel} key={channel.id} />);
    return <div className="channel-list">{listItems}</div>;
  }
}
ChannelList.propTypes = {
  channels: PropTypes.arrayOf(ChannelShape)
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
  error: PropTypes.instanceOf(Error)
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
Dashboard.propTypes = {
  channels: PropTypes.arrayOf(ChannelShape),
  error: PropTypes.instanceOf(Error)
};
