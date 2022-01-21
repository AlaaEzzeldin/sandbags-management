import prioritiesAPI from '../../api/priorities';

const state = {
  priorities: [],
}

const getters = {
  getPriorities(state) {
    return state.priorities;
  }
}

const actions = {
  loadPriorities({commit}) {
    prioritiesAPI.index()
      .then(function (response) {
        commit('SET_PRIORITIES', response.data);
      })
      .catch(function (error) {
        console.log(error);
      });
  }
}

const mutations = {
  SET_PRIORITIES(state, priorities) {
    state.priorities = priorities;
  },
}

export default {
  state,
  getters,
  actions,
  mutations
}
