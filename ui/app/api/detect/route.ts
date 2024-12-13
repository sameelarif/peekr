import { NextResponse } from "next/server";
import { detectAntibotProviders } from "@/lib/antibot-detector";

export async function POST(req: Request) {
  const { url } = await req.json();

  if (!url) {
    return NextResponse.json({ error: "URL is required" }, { status: 400 });
  }

  try {
    const providers = await detectAntibotProviders(url);
    return NextResponse.json({ providers });
  } catch (error) {
    return NextResponse.json(
      { error: "Failed to detect antibot providers" },
      { status: 500 }
    );
  }
}
