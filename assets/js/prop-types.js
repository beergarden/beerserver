/* @flow */
import { PropTypes } from 'react';

export const DatapointShape = PropTypes.shape({
  at: PropTypes.instanceOf(Date).isRequired,
  value: PropTypes.number.isRequired
});

export const ChannelShape = PropTypes.shape({
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  datapoints: PropTypes.arrayOf(DatapointShape).isRequired
});
