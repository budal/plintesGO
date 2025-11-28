CREATE OR REPLACE FUNCTION compute_holiday(
    referenced_by_easter boolean,
    current_year integer,
    date timestamp with time zone,
    difference_str character varying DEFAULT '0 seconds'
)
RETURNS timestamp
LANGUAGE plpgsql
AS $$
DECLARE
    diff interval;
    result timestamp;
BEGIN
    BEGIN
        diff := difference_str::interval;
    EXCEPTION WHEN others THEN
        diff := '0 seconds'::interval;
    END;

    result := CASE
        WHEN referenced_by_easter THEN
            (easter(current_year) + diff)
        ELSE
            make_timestamp(
                current_year,
                EXTRACT(MONTH   FROM date)::int,
                EXTRACT(DAY     FROM date)::int,
                EXTRACT(HOUR    FROM date)::int,
                EXTRACT(MINUTE  FROM date)::int,
                EXTRACT(SECOND  FROM date)::int
            )
    END;

    RETURN result;
END;
$$;
