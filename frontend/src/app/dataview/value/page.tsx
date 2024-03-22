"use client";

import ValuePage from '@/components/datacryp/value-page';
import { Suspense } from 'react';

export default function DataviewValuePage() {
  return (
    <Suspense>
      <ValuePage />
    </Suspense>
  );
}