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


export default({
    state,
    mutations,
    actions,
    getters
});
