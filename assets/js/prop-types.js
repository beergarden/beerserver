/* @flow */
import { PropTypes } from 'react';

export var DatapointShape = PropTypes.shape({
  at: PropTypes.instanceOf(Date).isRequired,
  value: PropTypes.number.isRequired
});

export var ChannelShape = PropTypes.shape({
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  datapoints: PropTypes.arrayOf(DatapointShape)
});
