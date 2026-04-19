import styles from "./AvatarDisplay.module.css";
import { ITEM_MASTER } from "../constants/items";

type Props = {
  hatId: number;
};

export const AvatarDisplay = ({ hatId }: Props) => {
  const item = ITEM_MASTER[hatId] || ITEM_MASTER[0];

  return (
    <div className={styles.container}>
      <div className={styles.imageBox}>
        <img src={item.img} alt={item.name} className={styles.image} />
      </div>
      <div>
        <p className={styles.label}>現在の装備</p>
        <p className={styles.itemName}>{item.name}</p>
      </div>
    </div>
  );
};