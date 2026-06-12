import './Cell.css'

export const Cell = ({ idx, value, onClick }) => {
  return <button className="cell" onClick={() => onClick(idx)}>{value}</button>;
};
