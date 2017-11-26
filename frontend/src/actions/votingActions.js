import * as types from '../constants/actionTypes';

import {getFormattedDateTime} from '../utils/dates';


export function upVoteSong(roomID, songID, userID) {
  return function(dispatch){

    const request = axios.post('http://127.0.0.1:8080/api/v1/room/${roomID}/songs/{songID}/upvote?user=${userID}')
    .then(function (response) {
      console.log('upvoted');
      dispatch({
          type: types.UPVOTE_SONG,
      });
    });
  }
}

export function downVoteSong(roomID, songID, userID) {
  return function(dispatch){

    const request = axios.post('http://127.0.0.1:8080/api/v1/room/${roomID}/songs/{songID}/downvote?user=${userID}')
    .then(function (response) {
      console.log('downvoted');
      dispatch({
          type: types.DOWNVOTE_SONG,
      });
    });
  }
}
