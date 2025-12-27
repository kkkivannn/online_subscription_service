create
extension if not exists "uuid-ossp"; -- расширение для генерации UUID

create table services
(
    id         uuid primary key default uuid_generate_v4(),  -- уникальный идентификатор записи
    name       text         not null,                             -- название сервиса
    price      integer      not null,                             -- стоимость месячной подписки в рублях
    user_id    uuid         not null,                             -- ID пользователя (UUID)
    start_date timestamp    not null                              -- дата начала подписки (месяц и год)
        check (start_date = date_trunc('month', start_date)),     -- всегда первый день месяца
    end_date   timestamp    null                                  -- дата конца подписки (месяц и год)
        check (start_date = date_trunc('month', end_date))
);
