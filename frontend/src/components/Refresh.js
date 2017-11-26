import React from 'react';
import PropTypes from 'prop-types';
import FuelSavingsResults from './FuelSavingsResults';
import FuelSavingsTextInput from './FuelSavingsTextInput';
import {fuelSavings} from '../types';

function Refresh({onRefreshClick}) {
  return (
    <div>
      <button type="submit" onClick={onRefreshClick}>Refresh</button>
    </div>
  );
}

// const { func } = PropTypes;

// FuelSavingsForm.propTypes = {
//   onRefreshClick: func.isRequired,
// };

export default Refresh;
