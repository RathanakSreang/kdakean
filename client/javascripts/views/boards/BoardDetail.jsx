import React, {Component} from 'react';
import {Switch, Route} from 'react-router-dom';
import { connect } from 'react-redux';
import { fetchBoard } from './../../redux/actions/board.actions.js';

@connect((store) => {
  return {
    ...store.boardReducers
  };
}, { fetchBoard })
class BoardDetail extends Component {
  componentWillMount() {
    const {params} = this.props.match;
    this.props.fetchBoard(params.slug);
  }
  componentDidUpdate (prevProps) {
    const {params} = this.props.match;
    const preParams = prevProps.match.params;
    if(preParams.slug !== params.slug) {
      this.props.fetchBoard(params.slug);
    }
  }

  render() {
    const {match, board} = this.props;
    return (
      <div className="container board-detail">
        <h1>{board.name}</h1>
        <div>Board navbar</div>
        <Switch>
          <Route exact path={`${match.url}/t/:task_slug`} render={(props) => <TaskDetail />} />
        </Switch>
      </div>
    );
  }
}

export default BoardDetail;
