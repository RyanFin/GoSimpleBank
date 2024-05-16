-- Reverse the code written in the migrate up script
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

-- retreive foreign key constraint name by clicking on 'info' in TablePlus
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";


-- Drop the users table
DROP TABLE IF EXISTS "users";