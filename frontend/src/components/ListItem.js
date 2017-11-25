import React from 'react';
import PropTypes from 'prop-types';
import '../styles/list-page.css';



function ListItem({index}) {

  return (
    <table className='tr-zone'>
      <tbody>
        <tr >
          <td className="fuel-savings-label"><button className='upButton'></button></td>
          <td className='song'> Song: {index} </td>
          <td className="dbutt"><button className='downButton'></button></td>
      </tr>
      </tbody>
    </table>
  );
};

ListItem.propTypes = {
  index: PropTypes.string.isRequired
};

export default ListItem;