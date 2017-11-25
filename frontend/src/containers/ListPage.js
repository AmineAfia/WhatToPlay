import React, {Component} from 'react';
import {render} from 'react-dom';
import ReactList from 'react-list';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as actions from '../actions/songsRefreshActions';
import '../styles/list-page.css';
import ListItem from '../components/ListItem';

const renderItem = (index, key) =>
  <div key={key} className={'item' + (index % 2 ? '' : ' even')}>
    {index}
  </div>;
renderItem.toJSON = () => renderItem.toString();

const getHeight = index => 50; // + (10 * (index % 10));
getHeight.toJSON = () => getHeight.toString();

const getWidth = index => 100 + (10 * (index % 10));
getWidth.toJSON = () => getWidth.toString();

const renderVariableHeightItem = (index, key) =>
  <div
    key={key}
    className={'item' + (index % 2 ? '' : ' even')}
    style={{lineHeight: `${getHeight(index)}px`}}
  >
    <ListItem index={index.toString()} />
  </div>;
renderVariableHeightItem.toJSON = () => renderVariableHeightItem.toString();


const examples = [
  // {
  //   length: 10,
  //   itemRenderer: renderVariableHeightItem
  // },
  // {
  //   length: 10,
  //   itemRenderer: renderVariableHeightItem,
  //   type: 'variable'
  // },
  // {
  //   length: 10,
  //   itemRenderer: renderVariableHeightItem,
  //   itemSizeGetter: getHeight,
  //   type: 'variable'
  // },
  {
    length: 1000,
    initialIndex: 10,
    itemRenderer: renderVariableHeightItem,
    itemSizeGetter: getHeight,
    type: 'variable'
  },
];

class ListPage extends Component {

  componentWillMount () {
    this.props.actions.getSongs();
  };

  renderExamples() {
    return examples.map((props, key) =>
      <div key={key} className={`example axis-${props.axis}`}>
       <div className='c_title'> <strong >Component</strong> </div>
        <div className='component'><ReactList {...props} /></div>
      </div>
    );
  }

  render() {
    return (
      <div className='example'>
        <div className='header'>Songs list</div>
        <div className='examples'>{this.renderExamples()}</div>
        <div>{this.state.songs}</div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    songs: state.songsReducer
  };
}

function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators(actions, dispatch)
  };
}

export default connect(null, mapDispatchToProps)(ListPage);
//export default ListPage;
