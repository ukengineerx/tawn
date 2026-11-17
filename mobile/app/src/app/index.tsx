import { useEffect, useState } from "react";
import { StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";

import { colors } from "@/theme/tokens";

const apiURL = process.env.EXPO_PUBLIC_API_URL ?? "http://localhost:8080";

export default function HomeScreen() {
  const [status, setStatus] = useState("Checking the API…");

  useEffect(() => {
    fetch(`${apiURL}/health`)
      .then((response) => setStatus(response.ok ? "API is up" : "API returned an error"))
      .catch(() => setStatus("API is not reachable from this device"));
  }, []);

  return (
    <SafeAreaView style={styles.screen}>
      <Text style={styles.kicker}>Tawn</Text>
      <Text style={styles.title}>Home</Text>
      <View style={styles.card}>
        <Text style={styles.body}>{status}</Text>
        <Text style={styles.meta}>{apiURL}</Text>
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  screen: {
    flex: 1,
    backgroundColor: colors.background,
    paddingHorizontal: 24,
    paddingTop: 32,
  },
  kicker: {
    color: colors.muted,
    fontSize: 14,
    fontWeight: "600",
  },
  title: {
    marginTop: 8,
    color: colors.text,
    fontSize: 32,
    fontWeight: "600",
  },
  card: {
    marginTop: 28,
    backgroundColor: colors.surface,
    borderColor: colors.line,
    borderWidth: 1,
    borderRadius: 16,
    padding: 16,
    gap: 8,
  },
  body: {
    color: colors.text,
    fontSize: 16,
  },
  meta: {
    color: colors.muted,
    fontSize: 13,
  },
});
