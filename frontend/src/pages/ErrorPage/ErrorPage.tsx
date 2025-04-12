import { Box, Container, Title, Text, Button } from "@mantine/core";
import { useNavigate } from "react-router-dom";
import { IconHome, IconArrowLeft } from "@tabler/icons-react";
import styles from "./ErrorPage.module.css";
import animations from "@/styles/animations.module.css";

const ErrorPage = () => {
  const navigate = useNavigate();

  return (
    <Box className={styles.container}>
      <Container size="sm" p={0}>
        <Box className={styles.card}>
          {/* 404 Badge */}
          <div className={`${styles.badge} ${animations.pulsingAnimation}`}>
            <Text fw={700} size="24px" c="white">
              404
            </Text>
          </div>

          <Box mt={32}>
            <Title order={1} className={styles.title}>
              Oops! Page Not Found
            </Title>
            <Text c="dimmed" size="md" maw={360} mx="auto" mb={32}>
              The page you're looking for seems to have wandered off into the digital wilderness.
            </Text>

            {/* Animated illustration */}
            <Box className={`${styles.illustration} ${animations.floatingAnimation}`}>
              <svg width="160" height="80" viewBox="0 0 160 80" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M20,40 Q80,80 140,40" stroke="#e9ecef" strokeWidth="3" fill="none"/>
                <circle cx="20" cy="40" r="6" fill="#FF8A00"/>
                <circle cx="140" cy="40" r="6" fill="#FF4F00"/>
                <path d="M65,35 Q80,15 95,35" stroke="#FF8A00" strokeWidth="2" fill="none"/>
              </svg>
            </Box>

            <div className={styles.buttonContainer}>
              <Button
                variant="subtle"
                color="gray"
                size="md"
                px="xl"
                leftSection={<IconArrowLeft size={18} />}
                onClick={() => navigate(-1)}
                className={`${styles.button} ${styles.buttonBack}`}
              >
                Go Back
              </Button>
              <Button
                variant="filled"
                color="orange"
                size="md"
                px="xl"
                leftSection={<IconHome size={18} />}
                onClick={() => navigate("/")}
                className={`${styles.button} ${styles.buttonHome}`}
              >
                Return Home
              </Button>
            </div>
          </Box>
        </Box>
      </Container>
    </Box>
  );
};

export default ErrorPage; 