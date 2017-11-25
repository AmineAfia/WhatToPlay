import * as types from '../constants/actionTypes';

import {getFormattedDateTime} from '../utils/dates';


export function upVoteSong(songID) {
  return function(dispatch){

    const request = axios.post('http://127.0.0.1:8080/api/v1/room/27029bbe-7fd6-4512-846f-7763fc97dad7/songs/{songID}/upvote')
    .then(function (response) {
      console.log('upvoted');
      dispatch({
          type: types.UPVOTE_SONG,
      });
    });
  }
}

export function downVoteSong(songID) {
  return function(dispatch){

    const request = axios.post('http://127.0.0.1:8080/api/v1/room/27029bbe-7fd6-4512-846f-7763fc97dad7/songs/{songID}/downvote')
    .then(function (response) {
      console.log('downvoted');
      dispatch({
          type: types.DOWNVOTE_SONG,
      });
    });
  }
}
