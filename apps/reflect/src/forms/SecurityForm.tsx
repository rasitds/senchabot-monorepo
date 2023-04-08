import React from "react";
import { useSession } from "next-auth/react";
import { Stack, Typography, Divider, Button } from "@mui/material";
import LinkAccount from "../components/auth/LinkAccount";
import FormTitle from "../components/FormTitle";
import { capitalizeWord } from "../utils/functions";
import { trpc } from "../utils/trpc";

const SecurityForm = () => {
  const { data: session } = useSession();
  const email = session?.user?.email || null || undefined;
  const [isLoading, setIsLoading] = React.useState(true);
  const accounts = trpc.security.getAccounts.useQuery();

  const [showEmailAddress, setShowEmailAddress] = React.useState(false);

  React.useEffect(() => {
    if (!accounts.isLoading) setIsLoading(false);
  }, [accounts.isLoading]);

  return (
    <>
      <FormTitle titleText="Link Account" />

      <Stack
        spacing={2}
        direction="row"
        sx={{ display: { xs: "none", md: "flex" } }}>
        <LinkAccount />
      </Stack>

      <Stack direction="column" sx={{ display: { xs: "flex", md: "none" } }}>
        <LinkAccount />
      </Stack>

      <Stack spacing={0.5} direction="column" sx={{ mt: 2 }}>
        <Typography>
          Linked accounts:{" "}
          {isLoading
            ? "Loading..."
            : accounts.data?.map(
                (account, index) =>
                  accounts &&
                  (accounts.data && accounts.data.length - 1 === index
                    ? (index > 2 ? "and " : "") +
                      capitalizeWord(account.provider)
                    : capitalizeWord(account.provider) +
                      (index < 2 ? " and " : ", ")),
              )}
        </Typography>
        <Typography>
          Your e-mail address:{" "}
          {showEmailAddress
            ? email && email.substring(email?.length / 2, email?.length)
            : email &&
              email
                .substring(email?.length / 2, email?.length)
                .replace(/./g, "*")}
        </Typography>
        <Button
          sx={{
            mt: 1,
            width: "fit-content",
            "@media (max-width: 600px)": {
              width: "100%",
            },
          }}
          onClick={() => setShowEmailAddress(!showEmailAddress)}
          variant="outlined">
          Show/Hide
        </Button>
      </Stack>

      <Divider
        orientation="horizontal"
        flexItem
        sx={{ mt: 2, mb: 2, width: "100%" }}
      />
    </>
  );
};

export default SecurityForm;