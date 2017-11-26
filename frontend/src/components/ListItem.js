import React from 'react';
import PropTypes from 'prop-types';
import '../styles/list-page.css';

function ListItem({index, onUpvoteClick, onDownvoteClick}) {

  return (
    <table className='tr-zone'>
      <tbody>
        <tr>
          <td className="fuel-savings-label"><button className='upButton' type="submit" onClick={onUpvoteClick}></button></td>
          <td className='song'> {index} </td>
          <td className="dbutt"><button className='downButton' type="submit" onClick={onDownvoteClick}></button></td>
      </tr>
      </tbody>
    </table>
  );
};

ListItem.propTypes = {
  index: PropTypes.string.isRequired
};

export default ListItem;