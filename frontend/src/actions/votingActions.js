import * as types from '../constants/actionTypes';

import {getFormattedDateTime} from '../utils/dates';


export function upvoteSong() {
  return {
    type: types.UPVOTE_SONG,
    dateModified: getFormattedDateTime(),
  };
}
