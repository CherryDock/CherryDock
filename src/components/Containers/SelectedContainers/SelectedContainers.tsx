import React from 'react';
import style from './selected-containers.module.scss';
import cx from 'classnames';

interface SelectedContainersProps {
    numContainers: number;
    deselectedContainer(): void;
}

function SelectedContainers({ numContainers, deselectedContainer }: SelectedContainersProps) {
    return (
        <div className={style.containerSelection}>
            <span className={style.label}>Selected Containers ({numContainers})</span>
            <button className={style.button} onClick={deselectedContainer}>
                <i className={cx("fas fa-times", style.closeIcon)}></i>
                <span className={style.closeLabel}>Deselect</span>
            </button>
        </div>
    )
}

export default SelectedContainers
