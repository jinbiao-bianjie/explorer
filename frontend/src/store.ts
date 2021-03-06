import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
const store = new Vuex.Store({
  state:{
    flShowSelectOption: false,
    currentEnv:'',
    isMobile:false,
    validatorTabIndex:0,
    flShowIpt:false,
    currentParticipationValue:"",
    currentYesValue:'',
    currentNoWithVetoValue:'',
    currentAbstainValue: '',
    currentNoValue: '',
    currentSkinStyle:'default'
  },
  mutations:{
    currentNoValue(state,data){
      state.currentNoValue = data
    },
    currentAbstainValue(state,data){
      state.currentAbstainValue = data
    },
    currentNoWithVetoValue(state,data){
      state.currentNoWithVetoValue = data
    },
    currentParticipationValue(state,data){
      state.currentParticipationValue = data
    },
    currentYesValue(state,data){
      state.currentYesValue = data
    },
    flShowSelectOption(state,data){
      state.flShowSelectOption = data
    },
    currentEnv(state,data){
      state.currentEnv = data
    },
    isMobile(state,data){
      state.isMobile = data
    },
    validatorTabIndex(state,data){
      state.validatorTabIndex = data
    },
    flShowIpt(state,data){
      state.flShowIpt = data
    },
    currentSkinStyle(state,data){
      state.currentSkinStyle = data
    }
  }
})
export default store
