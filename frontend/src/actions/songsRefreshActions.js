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

export function getRooms() {
    const request = axios.get('http://127.0.0.1:8080/api/v1/room');
    console.log(request.data);
    
    return {
        type: types.GET_ROOMS,
        payload: request
    };
}

export function getSongs() {
    const request = axios.get('http://127.0.0.1:8080/api/v1/room/27029bbe-7fd6-4512-846f-7763fc97dad7')
    .then(function (response) {
      console.log(response.data.songs);
      return {
          type: types.GET_SONGS,
          payload: response
      };
    });
    return () => request;
}
