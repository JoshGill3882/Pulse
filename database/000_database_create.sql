-- Create a dedicated schema for application tables if it does not already exist.
CREATE SCHEMA IF NOT EXISTS pulse;

-- Ensure objects in this script (and others that set the search path) target the pulse schema.
SET search_path TO pulse, public;
