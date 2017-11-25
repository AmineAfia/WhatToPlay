import * as types from '../constants/actionTypes';
import axios from 'axios';

// export function getSongs() {
//   return function (dispatch) {
//       axios.get('/room')
//       .then(function (response) {
//           console.log(response);
//       })
//       .catch(function (error) {
//           console.log(error);
//       });
//       console.log('get room');
//     return dispatch({
//       type: types.GET_SONGS,
//       // dateModified: getFormattedDateTime(),
//       // settings
//     });
//   };
// };

export function getSongs() {
    const request = axios.get(`https://127.0.0.1/room`);
    console.log(request);
    
    return {
        type: types.GET_SONGS,
        payload: request
    };
}
