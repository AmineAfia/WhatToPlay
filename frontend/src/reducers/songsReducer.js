import {GET_SONGS, GET_UPDATE, GET_ROOMS} from '../constants/actionTypes';
import {necessaryDataIsProvidedToCalculateSavings, calculateSavings} from '../utils/fuelSavings';
import objectAssign from 'object-assign';
import initialState from './initialState';

// IMPORTANT: Note that with Redux, state should NEVER be changed.
// State is considered immutable. Instead,
// create a copy of the state passed and set new values on the copy.
// Note that I'm using Objectassign to create a copy of current state
// and update values on the copy.
export default function songsReducer(state = initialState.songs, action) {

  switch (action.type) {

    case GET_SONGS:
        // console.log(action.payload.data.songs);
        return objectAssign({}, state, {songs: action.payload.data});

    case GET_UPDATE:
        // console.log(action.payload.data.songs);
        return objectAssign({}, state, {songs: action.payload.data});

    default:
        return state;
  }
}
