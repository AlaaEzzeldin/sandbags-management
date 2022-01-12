import statisticschartAPI from "../../api/statisticschart";

const state = {
    statisticschart: [],
}

const getters = {
    getStatisticschart(state) {

        console.log(state.statisticschart);
        return state.statisticschart;
    },

}

const mutations = {
    SET_STATISTICSCHART(state, statisticschart) {
        state.statisticschart = statisticschart;
    }
};

const actions = {
    loadStatisticschart({commit}) {
        statisticschartAPI.index()
            .then(function (response) {
                commit('SET_STATISTICSCHART', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    }
};

<<<<<<< HEAD
=======

>>>>>>> 285d454f5d7d37370b314a62a759391e4489de7c
export default({
    state,
    mutations,
    actions,
    getters
});
