export async function detectAntibotProviders(url: string): Promise<string[]> {
  const providers: string[] = [];

  try {
    const response = await fetch(url, {
      headers: {
        "User-Agent":
          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
      },
    });

    const headers = response.headers;
    const body = await response.text();

    // Check for common antibot providers
    if (headers.get("server") === "cloudflare") {
      providers.push("Cloudflare");
    }

    if (body.includes("PerimeterX") || body.includes("px-captcha")) {
      providers.push("PerimeterX");
    }

    if (body.includes("DataDome")) {
      providers.push("DataDome");
    }

    if (
      body.includes("Akamai Bot Manager") ||
      headers.get("x-akamai-transformed")
    ) {
      providers.push("Akamai Bot Manager");
    }

    if (body.includes("hCaptcha")) {
      providers.push("hCaptcha");
    }

    if (body.includes("reCAPTCHA")) {
      providers.push("reCAPTCHA");
    }

    // Add more detection logic for other providers as needed

    return providers;
  } catch (error) {
    console.error("Error detecting antibot providers:", error);
    throw error;
  }
}
