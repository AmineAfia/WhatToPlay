import React from 'react';
import PropTypes from 'prop-types';

const ListItem = ({index}) => {
  return (
    <table>
      <tbody>
        <tr>
          <td className="fuel-savings-label"><button>Up</button></td>
          <td> Song: {index} </td>
          <td><button>Down</button></td>
      </tr>
      </tbody>
    </table>
  );
};

ListItem.propTypes = {
  index: PropTypes.object.isRequired
};

export default ListItem;