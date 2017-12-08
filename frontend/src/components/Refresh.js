import React from 'react';

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
