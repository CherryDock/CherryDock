import React from 'react';
import style from './container-card.module.scss';
import cx from 'classnames';

interface ContainerCardProps {
    id: number;
    containerName: string;
    containerState: string;
    imageName: string;
    launchedDate: string;
    selected: boolean;
    toggleContainerFunc(id: number): void;
}

function ContainerCard({
    id,
    containerName,
    containerState,
    imageName,
    launchedDate,
    selected,
    toggleContainerFunc }: ContainerCardProps) {

    function onCardClick() {
        toggleContainerFunc(id);
    }

    function onSeeMoreClick(event: React.MouseEvent<HTMLElement>) {
        event.stopPropagation();
        alert('Go to the details page of the container ' + id);
    }

    const cardClass = selected? [style.card, style.cardSelected] : [style.card];

    return (
        <div className={cx(cardClass)}
            onClick={onCardClick}
        >
            <div className={style.header}>
                <span className={style.containerName}>{containerName}</span>
                <span className={style.containerState}>{containerState}</span>
            </div>

            <div className={style.imageName}>
                <i className={cx("fas fa-image", style.icon)}></i>
                <span className={style.label}>{imageName}</span>
            </div>


            <div className={style.cardContent}>
                <div className={style.leftPart}>
                    <div className={style.more}>
                        <i className={cx("fas fa-info-circle", style.icon)}></i>
                        <button onClick={(event) => onSeeMoreClick(event)} className={style.button}>More info</button>
                    </div>
                </div>

                <div className={style.rightPart}>
                    <div className={style.dateContainer}>
                        <i className={cx("far fa-calendar-alt", style.icon)}></i>
                        <div className={style.dateValue}>
                            <span className={style.label}>Last launched</span>
                            <span className={style.date}>{launchedDate}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default ContainerCard
