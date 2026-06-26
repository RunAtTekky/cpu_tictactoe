import './Cell.css'

export const Cell = ({ idx, value, onClick, color }) => {
  return <button style={{color: color}} className="cell" onClick={() => onClick(idx)}>{value}</button>;
};
