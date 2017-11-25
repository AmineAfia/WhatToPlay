import React, {Component} from 'react';
import {render} from 'react-dom';
import ReactList from 'react-list';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as actions from '../actions/votingActions';
import '../styles/list-page.css';
import ListItem from '../components/ListItem';

const renderItem = (index, key) =>
  <div key={key} className={'item' + (index % 2 ? '' : ' even')}>
    {index}
  </div>;
renderItem.toJSON = () => renderItem.toString();

const getHeight = index => 30 + (10 * (index % 10));
getHeight.toJSON = () => getHeight.toString();

const getWidth = index => 100 + (10 * (index % 10));
getWidth.toJSON = () => getWidth.toString();

const renderVariableHeightItem = (index, key) =>
  <div
    key={key}
    className={'item' + (index % 2 ? '' : ' even')}
    style={{lineHeight: `${getHeight(index)}px`}}
  >
    <ListItem index={index} />
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
  saveFuelSavings = () => {
    this.props.actions.upvoteSong();
  }

  renderExamples() {
    return examples.map((props, key) =>
      <div key={key} className={`example axis-${props.axis}`}>
        <strong>Props</strong>
        <pre className='props'>{JSON.stringify(props, null, 2)}</pre>
        <strong>Component</strong>
        <div className='component'><ReactList {...props} /></div>
      </div>
    );
  }

  render() {
    return (
      <div className='example'>
        <div className='header'>Songs list</div>
        <div className='examples'>{this.renderExamples()}</div>
      </div>
    );
  }
}

function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators(actions, dispatch)
  };
}

//export default connect(mapDispatchToProps)(ListPage);
export default ListPage;
