CREATE TABLE "otp" (
  "id" serial,
  "otp" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "expires_at" timestamp NOT NULL,
  CONSTRAINT "otp_pkey" PRIMARY KEY ("id")
);
