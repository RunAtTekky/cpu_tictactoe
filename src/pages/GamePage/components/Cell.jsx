import './Cell.css'

export const Cell = ({ value, onClick }) => {
  return <button className="cell" onClick={onClick}>{value}</button>;
};
