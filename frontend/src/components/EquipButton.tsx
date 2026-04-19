import styles from "./EquipButton.module.css";

type Props = {
  itemId: number;
  label: string;
  onClick: (itemId: number) => void;
};

export const EquipButton = ({ itemId, label, onClick }: Props) => {
  return (
    <button className={styles.button} onClick={() => onClick(itemId)}>
      {label}
    </button>
  );
};