import React from 'react';
import PropTypes from 'prop-types';
import '../styles/list-page.css';
const ListItem = ({index}) => {
  return (
    <table>
      <tbody>
        <tr className='tr-zone'>
          <td className="fuel-savings-label"><button className='upButton'></button></td>
          <td> Song: {index} </td>
          <td><button className='downButton'></button></td>
      </tr>
      </tbody>
    </table>
  );
};

ListItem.propTypes = {
  index: PropTypes.object.isRequired
};

export default ListItem;