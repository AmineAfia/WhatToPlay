import React, {Component} from 'react';
import {render} from 'react-dom';
import ReactList from 'react-list';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import * as actions from '../actions/songsRefreshActions';
import '../styles/list-page.css';
import ListItem from '../components/ListItem';
import Refresh from '../components/Refresh';

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


// renderSongsItem() {
//   return (
//     this.props.songs.map(song, key){
//     <div
//       key={key}
//       className={'item' + (index % 2 ? '' : ' even')}
//       style={{lineHeight: `${getHeight(index)}px`}}
//     >
//       <ListItem index={song.toString()} />
//     </div>
//   });
// }
// const renderListItem = renderSongsItem();
// renderListItem.toJSON = () => renderListItem.toString();


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
  }
];

class ListPage extends Component {

  componentWillMount () {
    this.props.actions.getSongs(this.getUrlHash());
  };

  getUrlHash() {
    return this.props.location.pathname.split('/room/')[1];
  };

  // renderExamples() {
  //   return examples.map((props, key) =>
  //     <div key={key} className={`example axis-${props.axis}`}>
  //      <div className='c_title'> <strong >Component</strong> </div>
  //       <div className='component'><ReactList {...props} /></div>
  //     </div>
  //   );
  // }

  renderWidgets(songslist) {
      // const songslist = this.props.songs.songs || [];
      let v
      let songsArray = []
      for(v in songslist) {
        // console.log(songslist[v]);
        songsArray.push(songslist[v]);
      }
      return songsArray.map(function(song) {
          return (
          <div key={song.id}>
            <ListItem index={song.title.toString()} />
          </div>
          );
      });
  }

  // renderWidgetsUpdate() {
  //     const songslist = this.props.updates.updates || [];
  //     let v
  //     let songsArray = []
  //     for(v in songslist) {
  //       // console.log(songslist[v]);
  //       songsArray.push(songslist[v]);
  //     }
  //     return songsArray.map(function(song) {
  //         return (
  //         <div key={song.id}>
  //           <ListItem index={song.title.toString()} />
  //         </div>
  //         );
  //     });
  // }

  refreshList = () => {
    this.props.actions.refreshList(this.getUrlHash());
    console.log('start refresch');
    console.log(this.props.updates.songs);
    // console.log(this.props);
    this.renderWidgets(this.props.updates.songs);
    this.forceUpdate()
    console.log('end refresch');
  }

  render() {
    return (
      <div className='example'>
        <div className='header'>Songs list</div>
        <div className='component'>{this.renderWidgets(this.props.songs.songs)}</div>
        <div className='ref' ><Refresh className='refb' onRefreshClick={this.refreshList} /></div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    songs: state.songsReducer.songs,
    updates: state.songsReducer.updates
  };
}

function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators(actions, dispatch)
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(ListPage);
//export default ListPage;
