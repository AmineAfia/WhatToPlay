import { combineReducers } from 'redux';
import fuelSavings from './fuelSavingsReducer';
import songsReducer from './songsReducer';
import { routerReducer } from 'react-router-redux';

const rootReducer = combineReducers({
  fuelSavings,
  songsReducer,
  routing: routerReducer
});

export default rootReducer;
