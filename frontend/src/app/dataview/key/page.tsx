"use client";

import KeyPage from '@/components/datacryp/key-page';
import { Suspense } from 'react';

export default function DataviewKeyPage() {
  return (
    <Suspense>
      <KeyPage />
    </Suspense>
  );
}
