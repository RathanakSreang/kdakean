import { boardService } from '../_services';
import { boardConstants, modalConstants } from '../_constants';
import { history } from '../_helpers';

export function createBoard(boardParams) {
  return function(dispatch) {
    boardService.createBoard(boardParams)
      .then((response) => {
        dispatch({type: modalConstants.MODAL_BOARD_OFF});
        const {board} = response.data;
        history.push('/k/' + board.slug);
      })
      .catch((err) => {
        console.log(err)
      })
  }
}

export function updateBoard(boardParams) {
  return function(dispatch) {
    boardService.updateBoard(boardParams)
      .then((response) => {
        dispatch({type: boardConstants.UPDATE_BOARDS_SUCCESS, payload: response.data});
        dispatch({type: modalConstants.MODAL_BOARD_OFF});
      })
      .catch((err) => {
        console.log(err)
      })
  }
}

export function fetchBoard(boardSlug) {
  return function(dispatch) {
    boardService.getBoard(boardSlug, {})
      .then((response) => {
        dispatch({type: boardConstants.GET_BOARD_DETAIL_SUCCESS, payload: response.data});
      })
      .catch((err) => {
        console.log(err)
      })
  }
}

export function fetchBoards(options) {
  return function(dispatch) {
    boardService.getBoards(options)
      .then((response) => {
        dispatch({type: boardConstants.GET_BOARDS_SUCCESS, payload: response.data})
      })
      .catch((err) => {
      })
  }
}

export function deleteBoard(boardSlug) {
  return function(dispatch) {
    boardService.deleteBoard(boardSlug)
      .then((response) => {
        dispatch({type: boardConstants.DELETE_BOARDS_SUCCESS, payload: response.data});
      })
      .catch((err) => {
        console.log(err)
      })
  }
}
