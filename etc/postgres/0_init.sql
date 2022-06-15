
--- change / add / modify this schema as necessary
CREATE TABLE public.sample_table (
    id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    sample VARCHAR(255)
);