import { CUSTOM_ACTIONS } from "../../constants";

const declareResult = (has_won, setResult, getResultMessage) => {
  setResult(getResultMessage(has_won));
};

export const ACTION_HANDLERS = {
  [CUSTOM_ACTIONS.DECLARE_RESULT]: declareResult,
}
